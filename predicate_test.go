package funk

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
				value: "test",
				predicates: []func(string) bool{
					func(v string) bool { return strings.Contains(v, "est") },
					func(v string) bool { return len(v) < 10 },
					func(v string) bool { return len(v) > 2 },
				},
			},
			want: true,
		},
		{
			name: "Sanity int predicates",
			args: args{
				value: 4,
				predicates: []func(int) bool{
					func(v int) bool { return v < 5 },
					func(v int) bool { return v > 2 },
				},
			},
			want: true,
		},
		{
			name: "Failed predicate",
			args: args{
				value: "test",
				predicates: []func(string) bool{
					func(v string) bool { return strings.Contains(v, "est") },
					func(v string) bool { return len(v) > 10 },
					func(v string) bool { return len(v) > 2 },
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
				value: "test",
				predicates: []func(string) bool{
					func(v string) bool { return strings.Contains(v, "est") },
					func(v string) bool { return len(v) > 10 },
					func(v string) bool { return len(v) < 2 },
				},
			},
			want: true,
		},
		{
			name: "Sanity int predicates",
			args: args{
				value: 4,
				predicates: []func(int) bool{
					func(v int) bool { return v > 5 },
					func(v int) bool { return v > 2 },
				},
			},
			want: true,
		},
		{
			name: "All failed predicate",
			args: args{
				value: "test",
				predicates: []func(string) bool{
					func(v string) bool { return !strings.Contains(v, "est") },
					func(v string) bool { return len(v) > 10 },
					func(v string) bool { return len(v) < 2 },
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

func TestPredicatesImplPanics(t *testing.T) {
	type args struct {
		value        interface{}
		wantedAnswer bool
		predicates   interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "predicates are not collection",
			args: args{
				value:        nil,
				wantedAnswer: false,
				predicates:   nil,
			},
		},
		{
			name: "predicates are collection of strings",
			args: args{
				value:        nil,
				wantedAnswer: false,
				predicates:   []string{"hey"},
			},
		},
		{
			name: "predicate has 2 out parameters",
			args: args{
				value:        nil,
				wantedAnswer: false,
				predicates:   []func(string) (bool, error){ func(string) (bool, error){return true, nil}},
			},
		},
		{
			name: "predicate has out parameter of type string",
			args: args{
				value:        nil,
				wantedAnswer: false,
				predicates:   []func(string) string{ func(string) string{return ""}},
			},
		},
		{
			name: "predicate has 2 in parameters",
			args: args{
				value:        nil,
				wantedAnswer: false,
				predicates:   []func(string, bool) bool{ func(string, bool) bool{return true}},
			},
		},
		{
			name: "predicate has 0 in parameters",
			args: args{
				value:        nil,
				wantedAnswer: false,
				predicates:   []func() bool{ func() bool{return true}},
			},
		},
		{
			name: "value is not convertible to in parameter",
			args: args{
				value:        1,
				wantedAnswer: false,
				predicates:   []func(string) bool{ func(string) bool{return true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Panics(t, func() {predicatesImpl(tt.args.value, tt.args.wantedAnswer, tt.args.predicates)})
		})
	}
}
