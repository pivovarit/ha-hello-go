extern crate iron;
extern crate router;

use iron::prelude::*;
use iron::status;
use router::Router;

fn main() {
    let mut router = Router::new();
    println!("Hello, world!");

    router.get("/api/payments", handler, "/api/payments");

    Iron::new(router)
      .http("0.0.0.0:8080").unwrap();
}

fn handler(_: &mut Request) -> IronResult<Response> {
    println!("Got a request!");
    Ok(Response::with((status::Ok, "Payments!\n")))
}
