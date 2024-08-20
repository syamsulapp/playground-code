<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use App\Models\Course;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Validator;

class CourseController extends Controller
{
    public function CourseEnroll(Request $request)
    {
        $Validate = Validator::make($request->all(), [
            'title' => 'required|string|unique:courses,title',
            'description' => 'required|string',
            'total_videos' => 'required|integer',
        ]);
        if ($Validate->fails()) {
            return $this->Response($Validate, $Validate->errors(), 422, false);
        }

        $Employee = new Course(); //instance model
        $Employee->users_id = Auth::guard('api')->user()->id;
        $Employee->title = $request->title;
        $Employee->description = $request->description;
        $Employee->total_videos = $request->total_videos;
        $Employee->save();
        return $this->Response($Employee, 'Create Enroll Course');
    }

    public function ListCourse()
    {
        $Course = new Course();
        if (isset($Course)) {
            $ListCourse = $Course->all();
            return $this->Response($ListCourse, 'Data Course Di Temukan');
        } else {
            $ListCourse = 'Data course tidak di temukan';
            return $this->Response($Course, $ListCourse, 422, false);
        }
    }

    public function DeleteCourse($DeleteCourseID)
    {
        $Course = new Course();
        if (!$Course
            ->where('id', $DeleteCourseID)
            ->exists()) {
            return $this->Response($Course, 'No data course', 422, false);
        } else {
            $DeleteCourse = $Course->find($DeleteCourseID);
            $DeleteCourse->delete();
            return $this->Response($DeleteCourse, 'Berhasil Delete Course');
        }
    }
}
