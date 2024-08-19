<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Employee;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class ApiController extends Controller
{
    //add employee
    public function AddEmployee(Request $request)
    {
        $Validate = Validator::make($request->all(), [
            'name' => 'required|string|min:4',
            'email' => 'required|email|unique:employees,email',
            'phone_number' => 'required|string|min:12|unique:employees,phone_number',
            'age' => 'required|integer',
            'gender' => 'required|in:male,female,other',
        ]);
        if ($Validate->fails()) {
            return $this->Response($Validate, $Validate->errors(), 422, false);
        }

        $Employee = new Employee(); //instance model
        $Employee->name = $request->name;
        $Employee->email = $request->email;
        $Employee->phone_number = $request->phone_number;
        $Employee->age = $request->age;
        $Employee->gender = $request->gender;
        $Employee->save();
        return $this->Response($Employee, 'Berhasil Create Data Employee');
    }


    //list employee
    public function ListEmployee()
    {
        $Employee = new Employee();
        if (!empty($Employee)) {
            $ListEmployees = $Employee->all();
        } else {
            $ListEmployees = 'Data Employee tidak di temukan';
        }
        return $this->Response($ListEmployees, 'Data Employee Di Temukan');
    }

    //detail detail
    public function DetailEmployee($EmployeeID)
    {
        $Employee = new Employee();
        if (!$Employee->where('id', $EmployeeID)->exists()) {
            return $this->Response($Employee, 'No data employee', 422, false);
        } else {
            return $this->Response($Employee->find($EmployeeID), 'Found Employees');
        }
    }

    //update detail
    public function UpdateEmployee($EmployeeID, Request $request)
    {
        $Employee = new Employee();
        if (!$Employee->where('id', $EmployeeID)->exists()) {
            return $this->Response($Employee, 'No data employee', 422, false);
        } else {
            $UpdateEmployee = $Employee->find($EmployeeID);
            $UpdateEmployee->name = $request->name ? $request->name : $UpdateEmployee->name;
            $UpdateEmployee->email = $request->email ? $request->email : $UpdateEmployee->email;
            $UpdateEmployee->phone_number = $request->phone_number ? $request->phone_number : $UpdateEmployee->phone_number;
            $UpdateEmployee->age = $request->age ? $request->age : $UpdateEmployee->age;
            $UpdateEmployee->gender = $request->gender ? $request->gender : $UpdateEmployee->gender;
            $UpdateEmployee->save();
            return $this->Response($request->all(), 'Berhasil Update Data Employee');
        }
    }

    //delete detail
    public function DeleteEmployee($EmployeeID) {}
}
