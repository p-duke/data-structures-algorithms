package main

import (
	"fmt"
	"reflect"
)

type Subscription struct {
	// Fill In
}

type User struct {
	// Fill In
}

func monthlyCharge(month string, subscription Subscription, users []User) float64 {
	// Solution here
	return 0.00
}

func main() {
	tests := []struct {
		desc         string
		month        string
		subscription Subscription
		users        []User
		want         float64
	}{
		{
			desc:         "Single User Active for the Entire Month",
			month:        "2023-05",
			subscription: Subscription{MonthlyRate: 100.00},
			users: []User{
				{ActivatedAt: "2023-05-01", DeactivatedAt: "2023-05-31"},
			},
			want: 100.00,
		},
		{
			desc:         "Single User Active for Part of the Month",
			month:        "2023-05",
			subscription: Subscription{MonthlyRate: 200.00},
			users: []User{
				{ActivatedAt: "2023-05-10", DeactivatedAt: "2023-05-20"},
			},
			want: 71.00,
		},
		{
			desc:         "Multiple Users Active for Different Ranges",
			month:        "2023-05",
			subscription: Subscription{MonthlyRate: 300.00},
			users: []User{
				{ActivatedAt: "2023-05-01", DeactivatedAt: "2023-05-10"},
				{ActivatedAt: "2023-05-05", DeactivatedAt: "2023-05-20"},
				{ActivatedAt: "2023-05-15", DeactivatedAt: ""},
			},
			want: 290.33,
		},
		{
			desc:         "User Spanning Across Month Boundaries",
			month:        "2023-05",
			subscription: Subscription{MonthlyRate: 150.00},
			users: []User{
				{ActivatedAt: "2023-04-25", DeactivatedAt: "2023-05-05"},
				{ActivatedAt: "2023-05-15", DeactivatedAt: "2023-05-30"},
			},
			want: 100.00,
		},
		{
			desc:         "User with No Deactivation Date (Active Until Present)",
			month:        "2023-06",
			subscription: Subscription{MonthlyRate: 250.00},
			users: []User{
				{ActivatedAt: "2023-06-01", DeactivatedAt: ""},
			},
			want: 250.00,
		},
		{
			desc:         "Leap Year February",
			month:        "2023-02",
			subscription: Subscription{MonthlyRate: 100.00},
			users: []User{
				{ActivatedAt: "2024-02-01", DeactivatedAt: ""},
			},
			want: 100.00,
		},
		{
			desc:         "User Activated Mid-Month",
			month:        "2023-07",
			subscription: Subscription{MonthlyRate: 200.00},
			users: []User{
				{ActivatedAt: "2023-07-15", DeactivatedAt: ""},
			},
			want: 109.68,
		},
		{
			desc:         "Multiple Users with Overlapping Ranges",
			month:        "2023-11",
			subscription: Subscription{MonthlyRate: 400.00},
			users: []User{
				{ActivatedAt: "2023-11-01", DeactivatedAt: "2023-11-10"},
				{ActivatedAt: "2023-11-05", DeactivatedAt: "2023-11-15"},
				{ActivatedAt: "2023-11-20", DeactivatedAt: ""},
			},
			want: 400.00,
		},
		{
			desc:         "User Active for One Day",
			month:        "2023-08",
			subscription: Subscription{MonthlyRate: 120.00},
			users: []User{
				{ActivatedAt: "2023-08-15", DeactivatedAt: "2023-08-15"},
			},
			want: 3.87,
		},
		{
			desc:         "No Active Users",
			month:        "2023-09",
			subscription: Subscription{MonthlyRate: 200.00},
			users: []User{},
			want: 0.00,
		},
	}

	for _, tc := range tests {
		got := monthlyCharge(tc.month, tc.subscription, tc.users)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("%s: FAILED! got: %f want: %f\n", tc.desc, got, tc.want)
		} else {
			fmt.Printf("%s: PASSED!\n", tc.desc)
		}
	}
}
