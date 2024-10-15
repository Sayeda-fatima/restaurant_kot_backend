<?php

namespace App\Http\Controllers\API\Auth;

use Hash;
use App\Http\Controllers\Controller;
use App\Http\Requests\Auth\LoginRequest;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use App\Models\User;
use Illuminate\View\View;
use Illuminate\Support\Str;


class AuthenticatedSessionController extends Controller
{
    /**
     * Display the login view.
     */
    public function create(): View
    {
        return view('auth.login');
    }

    /**
     * Handle an incoming authentication request.
     */
    public function store(LoginRequest $request)
    { {
            $credentials = $request->only('email', 'password');
            $user = User::where('email', $credentials['email'])->first();
            if (!$user || !Hash::check($credentials['password'], $user->password)) {
                return response()->json(['error' => 'Invalid credentials'], 401);
            }
            $user->api_token = hash('sha256',Str::random(60));
            $user->save();
            return response()->json(['user' => $user, 'api_token' => $user->api_token], 200);
        }
    }

    /**
     * Destroy an authenticated session.
     */
    public function destroy(Request $request): RedirectResponse
    {
        Auth::guard('api')->logout();

        $request->session()->invalidate();

        $request->session()->regenerateToken();

        return redirect('/');
    }
}
