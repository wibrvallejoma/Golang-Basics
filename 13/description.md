# How to make packages public and Private

If a package has it's first letter on lower case it will Private.
e.g.
    someClass

Otherwise, if it's has the first letter as Capital letter it will Public.
e.g.
    SomeClass

Also, to make its properties as Public, they first letter must be Mayus.
e.g.
    ```go
    type Car struct {
        Brand string
        Year string
    }
    ```
And a comment above the struct must be added.

## Packages
For best practices the package folder and the file must have the same name.