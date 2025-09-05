# icfp2025

[ICFP Contest 2025](https://icfpcontest2025.github.io/)

At this point I'm planning on using `golang`.

`python` is ready for tooling and as a backup.

## [Task](https://icfpcontest2025.github.io/task.html)

### Quotes

Wir müssen wissen, wir wer-
den wissen: We must know.  We will know.

### Registration

site: <https://31pwr5t6ij.execute-api.eu-west-2.amazonaws.com>

```sh
$curl -X POST https://31pwr5t6ij.execute-api.eu-west-2.amazonaws.com/register -H "Content-Type: application/json" -d '{"name":"Slow", "pl":"USA", "email":"grahamsspiers@gmail.com"}'
{"id":"grahamsspiers@gmail.com ffrbCqDWn7ARkZ9pR26Frg"}
```

id: <grahamsspiers@gmail.com> ffrbCqDWn7ARkZ9pR26Frg

## Plan

- Read the Spec
- Register
        ID = ffrbCqDWn7ARkZ9pR26Frg
- Look at a problem

    ```sh
    $curl -X POST https://31pwr5t6ij.execute-api.eu-west-2.amazonaws.com/select -H "Content-Type: application/json" -d '{"id":"grahamsspiers@gmail.com ffrbCqDWn7ARkZ9pR26Frg", "problemName":"probatio"}'
    {"problemName":"probatio"}

    $curl -X POST https://31pwr5t6ij.execute-api.eu-west-2.amazonaws.com/explore -H "Content-Type: application/json" -d '{"id":"grahamsspiers@gmail.com ffrbCqDWn7ARkZ9pR26Frg", "plans":[0, 0, 0]}'
    {"results":[[0,0],[0,0],[0,0]],"queryCount":4}
    ```

- Come up with strategies
        Fill in *Strategies*.
- What do we need?
        Expand this plan.
- Write the module and tests
- Implement best strategy

## Strategies

We need to get to *size* rooms.

Strategy one is brute force.  I will work on this while I think of different
strategies.

## Implementation Notes

### Build

#### `python`

`python` utilty code.

#### makefile

The `makefile` should contain common targets.

### Module

The `aedificium` module should have parts for representation, communication and
solving.

I need to create tests with simple problems.

### Executables
