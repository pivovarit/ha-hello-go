extern crate iron;
#[macro_use]
extern crate mime;

use iron::prelude::*;
use iron::status;

fn main() {
    println!("Hello, world!");
    Iron::new(handler).http("localhost:8080").unwrap();
}

fn handler(_request: &mut Request) -> IronResult<Response> {
    println!("Got a request!");

    let mut response = Response::new();

    response.set_mut(status::Ok);
    response.set_mut(mime!(Text/Plain; Charset=Utf8));
    response.set_mut("Payments\n");
    Ok(response)
}
