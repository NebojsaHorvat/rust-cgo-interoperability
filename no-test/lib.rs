// Rust Library: lib.rs

use std::os::raw::c_int;

#[no_mangle]
pub extern "C" fn applyStateTransition(slot: c_int, get_state: extern "C" fn(c_int) -> c_int) -> c_int {
    println!("Function applyStateTransition from Rust called!");
    let state = get_state(slot);
    state + 1
}