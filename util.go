package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

// ActorName represents the name of an actor in the system
type ActorName string

const (
	ActorConfig    ActorName = "config"
	ActorFlow      ActorName = "flow"
	ActorLogger    ActorName = "logger"
	ActorRegistry  ActorName = "registry"
	ActorOperation ActorName = "operation"
)

// resolveIfNull resolves a field if it is zero/nil by sending a message to the actor with the given name and waiting for a response
// If the field is already set, it skips the resolution
func resolveIfNull[T any](fieldPtr *T, ctx actor.Context, actorName string, msg interface{}, timeout time.Duration) error {
	// Check if field is already set (non-zero/non-nil)
	val := reflect.ValueOf(fieldPtr).Elem()
	if !val.IsZero() {
		fmt.Fprintf(os.Stderr, "Field of type %s is already set, skipping resolution\n", reflect.TypeOf(fieldPtr).Elem().Name())
		// Already set, skip resolving
		return nil
	} else {
		fmt.Fprintf(os.Stderr, "Field of type %s is not set, resolving\n", reflect.TypeOf(fieldPtr).Elem().Name())
	}

	// Resolve if field is zero/nil
	pid := ctx.ActorSystem().NewLocalPID(actorName)
	future, err := ctx.RequestFuture(pid, msg, timeout).Result()
	if err != nil {
		return err
	}

	// Set the field with the resolved value
	result := future.(T)
	val.Set(reflect.ValueOf(result))

	fmt.Fprintf(os.Stderr, "Field %s resolved to %v\n", reflect.TypeOf(fieldPtr).Name(), result)
	return nil
}

// Syntactic sugar for resolving the config actor
func resolveConfig[T any](fieldPtr *T, ctx actor.Context) error {
	return resolveIfNull(fieldPtr, ctx, string(ActorConfig), GetConfig{}, 10*time.Second)
}

// resolveInterface resolves an interface field if it is nil by calling the factory function
// If the field is already set, it skips the resolution
func resolveInterface[T any](fieldPtr *T, factory func() T) {
	// Check if field is already set (non-nil)
	val := reflect.ValueOf(fieldPtr).Elem()
	if !val.IsZero() {
		fmt.Fprintf(os.Stderr, "Interface field %s is already set, skipping resolution\n", reflect.TypeOf(fieldPtr).Elem().Name())
		// Already set, skip resolving
		return
	}

	fmt.Fprintf(os.Stderr, "Interface field %s is not set, resolving\n", reflect.TypeOf(fieldPtr).Elem().Name())

	// Resolve by calling the factory function
	result := factory()
	val.Set(reflect.ValueOf(result))

	fmt.Fprintf(os.Stderr, "Interface field %s resolved\n", reflect.TypeOf(fieldPtr).Elem().Name())
}
