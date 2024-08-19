<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Employee extends Model
{
    protected $primarykey = 'id';
    protected $fillable = ['name', 'email', 'phone_number', 'age', 'gender'];

    use HasFactory;
}
