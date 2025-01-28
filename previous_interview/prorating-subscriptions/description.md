# Prorating Subscriptions

## Background
Our company has started selling to larger customers, so we are creating subscription tiers with different feature sets to cater to our customers needs. We previously charged every customer a flat fee per month, but now we plan on charging for the number of users active on the customers subscription plan. As a result, we're changing our billing system.

## Instructions
You've picked up the work item to implement the logic to compute the monthly charge.

We'd like you to implement a `monthlyCharge` function to calculate the total monthly bill for a customer.

Customers are billed based on their subscription tier. We charge them a prorated amount for the portion of the month each user's subscriptions was active. For example, if a user was activated or deactivated part way through the month, then we charge them only for the days their subscription was active.

We want to bill customers for all days users were active in the month (including any activation and deactivation dates, since the user had some access on those days).
- We do need to support historical calculations (previous dates)
- We only charge whole cents

### Notes
Here's an idea of how we might go about this:
1. Calculate a daily rate for the subscription tier
1. For each day of the month, identify which users has an active subscritpion on that day
1. Multiply the number of active users for the day by the daily rate to calculate the total for the day
1. Return the running total for the month at the end


## Testing
The provided unit tests only cover a few cases that one of your colleagues shared, so you should plan to add your own tests to ensure that your solution handles any edge cases
