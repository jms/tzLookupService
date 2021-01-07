package main

import "testing"

func Test_lookupTz(t *testing.T) {
	type args struct {
		Lat float32
		Lon float32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Managua Timezone", args{12, -86}, "America/Managua", false},
		{"Madrid Timezone", args{40.4381311, -3.8196196}, "Europe/Madrid", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lookupTz(tt.args.Lat, tt.args.Lon)
			if (err != nil) != tt.wantErr {
				t.Errorf("lookupTz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lookupTz() got = %v, want %v", got, tt.want)
			}
		})
	}
}
