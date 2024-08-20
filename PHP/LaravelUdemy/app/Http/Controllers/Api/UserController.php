<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use Tymon\JWTAuth\Facades\JWTAuth;

class UserController extends Controller
{
    public function Login(Request $request)
    {

        $Validate = Validator::make($request->all(), [
            'email' => 'required|email',
            'password' => 'required|string',
        ]);
        if ($Validate->fails()) {
            return $this->Response($Validate, $Validate->errors(), 422, false);
        }

        $Token = JWTAuth::attempt([
            'email' => $request->email,
            'password' => $request->password,
        ]);
        if (!empty($Token)) {
            return $this->Response($Token, 'Berhasil Login');
        } else {
            return $this->Response($Token, 'Email atau password salah', 422, false);
        }
    }

    public function Register(Request $request)
    {
        $Validate = Validator::make($request->all(), [
            'name' => 'required|string|min:4',
            'email' => 'required|email|unique:users,email',
            'phone_no' => 'required|string|unique:users,phone_no|min:12',
            'password' => 'required|string',
        ]);
        if ($Validate->fails()) {
            return $this->Response($Validate, $Validate->errors(), 422, false);
        }

        $Employee = new User(); //instance model
        $Employee->name = $request->name;
        $Employee->email = $request->email;
        $Employee->phone_no = $request->phone_no;
        $Employee->password = Hash::make($request->password);
        $Employee->save();
        return $this->Response($Employee, 'Berhasil Create User');
    }

    public function Logout() {}

    public function Profile()
    {
        return $this->Response(Auth::guard('api')->user(), 'Get Data Profile Successfully');
    }

    public function TokenRefresh() {}
}
