<?php

namespace Database\Factories;

use App\Models\Order;
use App\Models\Organization;
use App\Models\Product;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\OrderItem>
 */
class OrderItemFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'organization_id' => Organization::factory(),
            'order_id' => Order::factory(),
            'product_id' => Product::factory(),
            'product_quantity' => $this->faker->numberBetween(1, 30),
            'unit_product_price' => $this->faker->numberBetween(1,18),
            'tax' => $this->faker->numberBetween(1,18),
            'total_product_price' => $this->faker->numberBetween(2, 567)
        ];
    }
}
