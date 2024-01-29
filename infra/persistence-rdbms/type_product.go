package main

type Product struct {
	// gorm.Model
	ID        uint  `hera:"pk"`
	CreatedAt int64 `hera:"columnname:created_at"`
	UpdatedAt int64 `hera:"columnname:updated_at"`
	DeletedAt int64 `hera:"columnname:deleted_at"`

	Code  string
	Price uint
}
