<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;

class Controller extends BaseController
{
    use AuthorizesRequests, ValidatesRequests;


    protected static function Response($Data = 'No Data', string $Message = 'No Message', int $Code = 200,  $Status = true)
    {
        if ($Code != 200 || $Status != true) {
            $Res = array('status' => $Status, 'message' => $Message);
        } else {
            $Res = array('status' => $Status, 'message' => $Message, 'data' => $Data);
        }
        return response()->json($Res, $Code);
    }
}
