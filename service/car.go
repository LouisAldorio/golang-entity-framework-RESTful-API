package service

import (
	"context"
	"fmt"
	"log"
	"myapp/ent"
	"myapp/ent/car"
	"time"
	"myapp/ent/user"
)

func QueryUserCars(ctx context.Context, client *ent.Client, username string) ([]*ent.Car,error) {

	user := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name(username),
		).
		OnlyX(ctx)

	cars, err := user.QueryCars().WithCarUtilities().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user cars: %w", err)
	}

	// Query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed querying car %q owner: %w", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}

	return cars, nil
}

func QueryUserSpesificCar(ctx context.Context, a8m *ent.User, carModel string) (*ent.Car,error) {

	ford, err := a8m.QueryCars().
		Where(car.Model(carModel)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user cars: %w", err)
	}
	
	log.Println(ford)
	return ford, nil
}

func CreateCar(ctx context.Context, user *ent.User, client *ent.Client, carModel string) (*ent.Car, error) {

	car, err := client.Car.
		Create().
		SetModel(carModel).
		SetRegisteredAt(time.Now()).
		SetOwner(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", car)

	return car, nil
}
