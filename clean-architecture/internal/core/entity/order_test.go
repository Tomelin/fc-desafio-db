package entity

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestOrder_Validate(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		Stock       uint
		Amount      uint
		Price       float32
		Category    Category
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "FieldNil", wantErr: true},
		{name: "FieldEmpty", fields: fields{}, wantErr: true},
		{name: "FieldEmptyName", fields: fields{
			Name:        "",
			Description: "Golang",
			Stock:       1,
			Price:       5,
			Amount:      1,
			Category:    1,
		}, wantErr: true},
		{name: "FieldNameUndeclared", fields: fields{
			Description: "Golang",
			Stock:       1,
			Price:       5,
			Amount:      1,
			Category:    1,
		}, wantErr: true},
		{name: "FieldEmptyAmount", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       1,
			Amount:      0,
			Price:       5,
			Category:    1,
		}, wantErr: true},
		{name: "FieldAmountUndeclared", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       1,
			Price:       5,
			Category:    1,
		}, wantErr: true},
		{name: "FieldAmountGtStock", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       1,
			Amount:      10,
			Price:       5,
			Category:    1,
		}, wantErr: true},
		{name: "FieldAmountAndStockEqualZero", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       0,
			Amount:      0,
			Price:       5,
			Category:    1,
		}, wantErr: true},
		{name: "FieldAmountAndStockEqual", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       1,
			Amount:      1,
			Price:       5,
			Category:    1,
		}, wantErr: true},
		{name: "FieldPriceZero", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       1,
			Amount:      10,
			Price:       0,
			Category:    1,
		}, wantErr: true},
		{name: "FieldAllZero", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       0,
			Amount:      0,
			Price:       0,
			Category:    1,
		}, wantErr: true},
		{name: "FieldSuccessfully", fields: fields{
			Name:        "book",
			Description: "Golang",
			Stock:       10,
			Amount:      1,
			Price:       10.0,
			Category:    1,
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Stock:       tt.fields.Stock,
				Price:       tt.fields.Price,
				Category:    tt.fields.Category,
			}
			if err := o.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Order.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCategory_GetCategoryName(t *testing.T) {
	tests := []struct {
		name string
		c    Category
		want string
	}{
		{
			name: "Software",
			c:    1,
			want: "Software",
		},
		{
			name: "HomeAutomation",
			c:    2,
			want: "HomeAutomation",
		},
		{
			name: "BigData",
			c:    3,
			want: "BigData",
		},
		{
			name: "Unknow",
			c:    0,
			want: "Unknow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetCategoryName(); got != tt.want {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestOrderResponse_Validate(t *testing.T) {
	type fields struct {
		ID    string
		Order Order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "FieldsNil", wantErr: true},
		{name: "FieldsEmpty", fields: fields{}, wantErr: true},
		{name: "FieldsOrderError", fields: fields{
			ID: "0930474c-f672-4f8c-90af-2d13c16116fe",
			Order: Order{
				Name:        "",
				Description: "Golang",
				Stock:       1,
				Amount:      10,
				Price:       5,
				Category:    1,
			},
		}, wantErr: true},
		{name: "FieldsIDUndeclared", fields: fields{
			Order: Order{
				Name:        "book",
				Description: "Golang",
				Stock:       1,
				Amount:      10,
				Price:       5,
				Category:    1,
			},
		}, wantErr: true},
		{name: "FieldsIDEmpty", fields: fields{
			ID: "",
			Order: Order{
				Name:        "book",
				Description: "Golang",
				Stock:       1,
				Amount:      10,
				Price:       5,
				Category:    1,
			},
		}, wantErr: true},
		{name: "FieldsIDEmpty", fields: fields{
			ID: "0930474c-f672-4f8c-90af-2d13c16116fe",
			Order: Order{
				Name:        "book",
				Description: "Golang",
				Stock:       1,
				Amount:      10,
				Price:       5,
				Category:    1,
			},
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderResponse{
				ID:    tt.fields.ID,
				Order: tt.fields.Order,
			}
			if err := o.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
