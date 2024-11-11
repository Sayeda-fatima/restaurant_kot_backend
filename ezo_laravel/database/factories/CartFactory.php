<?php

namespace Database\Factories;

use App\Models\Organization;
use App\Models\Customer;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Cart>
 */
class CartFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'organization_id' => Organization::factory("id"),
            'customer_id' => Customer::factory(),
            'total_quantity' => $this->faker->numberBetween(1, 40)
        ];
    }
}
