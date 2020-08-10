package funk

import (
	"strings"
	"testing"
)

func TestAllPredicates(t *testing.T) {
	type args struct {
		value      interface{}
		predicates interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Sanity string predicates",
			args: args{
				value:      "test",
				predicates: []func(string)bool {
					func(v string) bool {return strings.Contains(v, "est")},
					func(v string) bool {return len(v) < 10},
					func(v string) bool {return len(v) > 2},
				},
			},
			want: true,
		},
		{
			name: "Sanity int predicates",
			args: args{
				value:      4,
				predicates: []func(int)bool {
					func(v int) bool {return v < 5},
					func(v int) bool {return v > 2},
				},
			},
			want: true,
		},
		{
			name: "Failed predicate",
			args: args{
				value:      "test",
				predicates: []func(string)bool {
					func(v string) bool {return strings.Contains(v, "est")},
					func(v string) bool {return len(v) > 10},
					func(v string) bool {return len(v) > 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllPredicates(tt.args.value, tt.args.predicates); got != tt.want {
				t.Errorf("AllPredicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyPredicates(t *testing.T) {
	type args struct {
		value      interface{}
		predicates interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Sanity string predicates",
			args: args{
				value:      "test",
				predicates: []func(string)bool {
					func(v string) bool {return strings.Contains(v, "est")},
					func(v string) bool {return len(v) > 10},
					func(v string) bool {return len(v) < 2},
				},
			},
			want: true,
		},
		{
			name: "Sanity int predicates",
			args: args{
				value:      4,
				predicates: []func(int)bool {
					func(v int) bool {return v > 5},
					func(v int) bool {return v > 2},
				},
			},
			want: true,
		},
		{
			name: "All failed predicate",
			args: args{
				value:      "test",
				predicates: []func(string)bool {
					func(v string) bool {return !strings.Contains(v, "est")},
					func(v string) bool {return len(v) > 10},
					func(v string) bool {return len(v) < 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyPredicates(tt.args.value, tt.args.predicates); got != tt.want {
				t.Errorf("AnyPredicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
