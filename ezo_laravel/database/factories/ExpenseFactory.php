<?php

namespace Database\Factories;

use App\Models\Supplier;
use App\Models\Organization;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Expense>
 */
class ExpenseFactory extends Factory
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
            'supplier_id' => Supplier::factory(),
            'supplier_name' => $this->faker->name(),
            'expense_category' => $this->faker->words('electricity', 'gas'),
            'total_amount' => $this->faker->numberBetween('100', '3000000000'),
            'amount_paid' => $this->faker->numberBetween('100', '3000000000'),
            'amount_due' => $this->faker->numberBetween('100', '3000000000'),
            'note' => $this->faker->text(),
            'mode_of_payment' => 'bank',//$this->faker->words('cash', 'bank')
        ];
    }
}
