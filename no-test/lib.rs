// Rust Library: lib.rs

#[no_mangle]
pub extern "C" fn applyStateTransition(slot: i32, get_state: extern "C" fn(i32) -> i32) -> i32 {
    println!("Function applyStateTransition from Rust called!");
    let state = get_state(slot);
    state + 1
}