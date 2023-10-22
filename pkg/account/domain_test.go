package account_test

import (
	"testing"

	"github.com/ageeknamedslickback/interest-rate/pkg/account"
)

func TestMapBalanceToApplicableInterest(t *testing.T) {
	type args struct {
		balance float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "balance first band(1000)",
			args:    args{balance: 999},
			want:    account.FirstBandInterest,
			wantErr: false,
		},
		{
			name:    "balance second band(5000)",
			args:    args{balance: 4999},
			want:    account.SecondBandInterest,
			wantErr: false,
		},
		{
			name:    "balance < third band(10000)",
			args:    args{balance: 9999},
			want:    account.ThirdBandInterest,
			wantErr: false,
		},
		{
			name:    "balance < last band(50000)",
			args:    args{balance: 49999},
			want:    account.LastBandInterest,
			wantErr: false,
		},
		{
			name:    "balance gte last band",
			args:    args{balance: 50000},
			want:    account.GTELastBandInterest,
			wantErr: false,
		},
		{
			name:    "negative balance",
			args:    args{balance: -50000},
			want:    account.NoOpInterest,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := account.MapBalanceToApplicableInterest(tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapBalanceToApplicableInterest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MapBalanceToApplicableInterest() = %v, want %v", got, tt.want)
			}
		})
	}
}
