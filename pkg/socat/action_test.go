package socat

import "testing"

func TestAddress_String(t *testing.T) {
	type fields struct {
		Address string
		Options []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Simple Test",
			fields: fields{
				Address: "UNIX-CONNECT:/tmp/foo",
				Options: []string{
					"fork",
					"reuseaddr",
				},
			},
			want: "UNIX-CONNECT:/tmp/foo,fork,reuseaddr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Address{
				Address: tt.fields.Address,
				Options: tt.fields.Options,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("Address.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
