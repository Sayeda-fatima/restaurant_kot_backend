<?php

namespace App\Policies;

use App\Models\ProductCategory;
use App\Models\User;
use Illuminate\Auth\Access\Response;

class ProductCategoryPolicy
{
    /**
     * Determine whether the user can view any models.
     */
    public function viewAny(User $user): bool
    {
        return $user->access_type === 'ADMIN' || $user->access_type === 'STAFF' || $user->access_type === 'SALES';
    }

    /**
     * Determine whether the user can view the model.
     */
    public function view(User $user): bool
    {
        return $user->access_type === 'ADMIN';
    }

    /**
     * Determine whether the user can create models.
     */
    public function create(User $user): bool
    {
        return $user->access_type === 'ADMIN' || $user->access_type === 'STAFF';
    }

    /**
     * Determine whether the user can update the model.
     */
    public function update(User $user, ?ProductCategory $productCategory): bool
    {
        return $user->access_type === 'ADMIN' || $user->access_type === 'STAFF';
    }

    /**
     * Determine whether the user can delete the model.
     */
    public function delete(User $user, ?ProductCategory $productCategory): bool
    {
        return $user->access_type === 'ADMIN' || $user->access_type === 'STAFF';
    }

    /**
     * Determine whether the user can restore the model.
     */
    public function restore(User $user, ProductCategory $productCategory): bool
    {
        return $user->access_type === 'ADMIN';
    }

    /**
     * Determine whether the user can permanently delete the model.
     */
    public function forceDelete(User $user, ProductCategory $productCategory): bool
    {
        return $user->access_type === 'ADMIN';
    }
}
