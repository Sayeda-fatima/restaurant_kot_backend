<?php

namespace Database\Factories;

use App\Models\Organization;
use App\Models\Cart;
use App\Models\Product;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\CartItem>
 */
class CartItemFactory extends Factory
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
            'cart_id' => Cart::factory("id"),
            'product_id' => Product::factory("id"),
            'product_quantity' => $this->faker->numberBetween(1,30)
        ];
    }
}
