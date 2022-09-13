package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/lander2k2/gqlgen-todos/graph/generated"
	"github.com/lander2k2/gqlgen-todos/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// create new todo
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		User:   &model.User{ID: input.UserID, Name: "user " + input.UserID},
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)

	fmt.Printf("%+v\n", r.TodoSubscribers)

	// notify subscribers
	r.mu.Lock()
	for _, subscriber := range r.TodoSubscribers {
		subscriber <- todo
	}
	r.mu.Unlock()

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// TodoNotifs is the resolver for the todoNotifs field.
func (r *subscriptionResolver) TodoNotifs(ctx context.Context) (<-chan *model.Todo, error) {
	id := randString(8)
	todos := make(chan *model.Todo, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.TodoSubscribers, id)
		r.mu.Unlock()
	}()
	r.mu.Lock()

	r.TodoSubscribers[id] = todos
	r.mu.Unlock()

	return todos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
