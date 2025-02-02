package sql

import "github.com/jimtrung/amazon/internal/database/migrations"

func SetupDatabase() error {
	// schema.go
	err := CreateSchema()
	if err != nil {
		return err
	}

	// 001_config_products.go
	err = migrations.BindDefaultProducts()
	if err != nil {
		return err
	}

	err = migrations.BindRuleProducts()
	if err != nil {
		return err
	}

	// 002_config_cart.go
	err = migrations.BindDefaultCart()
	if err != nil {
		return err
	}

	err = migrations.BindRuleCart()
	if err != nil {
		return err
	}

	// 003_add_to_cart.go
	err = migrations.CreateAddToCart()
	if err != nil {
		return err
	}

	// 004_delete_from_cart.go
	err = migrations.CreateDeleteFromCart()
	if err != nil {
		return err
	}

	// 005_update_cart.go
	err = migrations.CreateUpdateCart()
	if err != nil {
		return err
	}

	// 006_generate_uid.go
	err = migrations.CreateUIDGenerator()
	if err != nil {
		return err
	}

	// 007_create_user_database.go
	err = migrations.CreateTableUser()
	if err != nil {
		return err
	}

	// 008_config_users.go
	err = migrations.BindDefaultUsers()
	if err != nil {
		return err
	}

	// 008_config_users.go
	err = migrations.BindRuleUsers()
	if err != nil {
		return err
	}

	return nil
}
