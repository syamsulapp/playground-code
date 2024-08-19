<?php

use App\Http\Controllers\Api\ApiController;
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

//add employee
Route::get('add-employee', [ApiController::class, 'AddEmployee']);

//add employee
Route::get('update-employee', [ApiController::class, 'UpdateEmployee']);

//add employee
Route::get('delete-employee', [ApiController::class, 'DeleteEmployee']);
