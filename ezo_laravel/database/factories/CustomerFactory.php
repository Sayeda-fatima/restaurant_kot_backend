<?php

namespace Database\Factories;

use App\Models\Organization;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Customer>
 */
class CustomerFactory extends Factory
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
            'name' => $this->faker->name,
            'phone_no' => $this->faker->unique()->numberBetween(1000000, 99999999),
            'category' => 'retail',#$this->faker->words('wholesale', 'retail'),
            'billing_address' => fake()->address(),
            'billing_province' => fake()->streetName(),
            'billing_postal_code' => fake()->postcode(),
            'delivery_address' => fake()->address(),
            'delivery_province' => fake()->streetName(),
            'delivery_postal_code' => fake()->postcode(),
            'gst_number' => 'GST'. fake()->numberBetween(1000000, 99999999),
            'billing_term' => $this->faker->word(),
            'billing_type' => 'debit',//$this->faker->words('debit', 'credit'),
            'date_of_birth' => fake()->date(),
            'whatsapp_alert' => 'Y'//$this->faker->words('Y', 'N')
        ];
    }
}
