<?php

use App\Http\Controllers\Api\ApiController;
use App\Http\Controllers\Api\CourseController;
use App\Http\Controllers\Api\UserController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

//list employee
Route::get('list-employee', [ApiController::class, 'ListEmployee']);

//detail employee
Route::get('detail-employee/{EmployeeID}', [ApiController::class, 'DetailEmployee']);

//add employee
Route::post('add-employee', [ApiController::class, 'AddEmployee']);

//add employee
Route::put('update-employee/{EmployeeID}', [ApiController::class, 'UpdateEmployee']);

//add employee
Route::delete('delete-employee/{EmployeeID}', [ApiController::class, 'DeleteEmployee']);

Route::prefix('users')->group(function () {
    //user login,register
    Route::post('login', [UserController::class, 'Login']);
    Route::post('register', [UserController::class, 'Register']);

    Route::middleware('auth:api')->group(function () {
        //user logout,profile,refresh token
        Route::post('logout', [UserController::class, 'Logout']);
        Route::post('profile', [UserController::class, 'Profile']);
        Route::post('refresh-token', [UserController::class, 'TokenRefresh']);

        //enroll course by users
        Route::prefix('course')->group(function () {
            Route::get('/', [CourseController::class, 'ListCourse']);
            Route::post('enroll', [CourseController::class, 'CourseEnroll']);
            Route::post('delete/{DeleteCourseID}', [CourseController::class, 'DeleteCourse']);
        });
    });
});
