package paginate_test

// import (
// 	"reflect"
// 	"testing"

// 	gorm_paginations "github.com/anggitrestuu/go-rest-api/pkg/paginations"
// )

// // TestParseFilters tests the parseFilters function
// func TestParseFilters(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		filters string
// 		want    []gorm_paginations.Filter
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Empty filter string",
// 			filters: "",
// 			want:    nil,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Invalid filter string",
// 			filters: "invalid",
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name:    "Only Operator",
// 			filters: "$eq:name=john;$like:;",
// 			want: []gorm_paginations.Filter{
// 				{
// 					Value:      "name = 'john'",
// 					Comparator: "AND",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Nested Valid filter string",
// 			filters: "$eq:account.id=1;",
// 			want: []gorm_paginations.Filter{
// 				{
// 					Value:      "account.id = '1'",
// 					Comparator: "AND",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Valid filter string",
// 			filters: "$eq:name=john;$like:email=john;$in:age=20,30;$btw:age=20,30;$or:$like:color=red;",
// 			want: []gorm_paginations.Filter{
// 				{
// 					Value:      "name = 'john'",
// 					Comparator: "AND",
// 				},
// 				{
// 					Value:      "email LIKE '%john%'",
// 					Comparator: "AND",
// 				},
// 				{
// 					Value:      "age IN (20,30)",
// 					Comparator: "AND",
// 				},
// 				{
// 					Value:      "age BETWEEN 20 AND 30",
// 					Comparator: "AND",
// 				},
// 				{
// 					Value:      "color LIKE '%red%'",
// 					Comparator: "OR",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := gorm_paginations.ParseFilterFromString(tt.filters)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ParseFilterFromString() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ParseFilterFromString() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
