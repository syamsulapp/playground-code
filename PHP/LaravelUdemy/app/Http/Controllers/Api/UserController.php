<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;

class UserController extends Controller
{
    public function Login() {}

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

    public function Profile() {}

    public function TokenRefresh() {}
}
