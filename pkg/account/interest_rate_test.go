package account_test

import (
	"testing"

	"github.com/ageeknamedslickback/interest-rate/pkg/account"
	"github.com/shopspring/decimal"
)

func TestCalculateInterestApplicable(t *testing.T) {
	type args struct {
		balance float64
	}
	tests := []struct {
		name    string
		args    args
		want    decimal.Decimal
		wantErr bool
	}{
		{
			name:    "balance first band(1000)",
			args:    args{balance: 999},
			want:    decimal.NewFromFloat(9.99),
			wantErr: false,
		},
		{
			name:    "balance second band(5000)",
			args:    args{balance: 4999},
			want:    decimal.NewFromFloat(74.99),
			wantErr: false,
		},
		{
			name:    "balance < third band(10000)",
			args:    args{balance: 9999},
			want:    decimal.NewFromFloat(199.98),
			wantErr: false,
		},
		{
			name:    "balance < last band(50000)",
			args:    args{balance: 49999},
			want:    decimal.NewFromFloat(1249.98),
			wantErr: false,
		},
		{
			name:    "balance gte last band",
			args:    args{balance: 50000},
			want:    decimal.NewFromFloat(1500),
			wantErr: false,
		},
		{
			name:    "negative balance",
			args:    args{balance: -50000},
			want:    decimal.NewFromFloat(account.NoOpFloatValue),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := account.CalculateInterestApplicable(tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateInterestApplicable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("CalculateInterestApplicable() = %v, want %v", got, tt.want)
			}
		})
	}
}
