<?php

abstract class SolveProblem
{
    abstract public function groupByOwners(array $files);
}

class Problem extends SolveProblem
{
    public function groupByOwners(array $files)
    {
        foreach ($files as $key => $value) {
            return "key: {$key} => value: {$value}";
        }
    }
}

$classProblem = new Problem();

$files = array(
    "input" => "randy",
    "code" => "stan",
    "output" => "randy"
);

echo $classProblem->groupByOwners($files);

/**
 * ['randy' => ['Input.txt', 'Output.txt'], 'stan' => ['code.py']]
 */
