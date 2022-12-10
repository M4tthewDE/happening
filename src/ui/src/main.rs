use log::info;
use web_sys::HtmlInputElement;
use yew::prelude::*;

#[function_component]
fn App() -> Html {
    let input_node_ref = use_node_ref();

    let onclick = {
        let input_node_ref = input_node_ref.clone();
        Callback::from(move |_| {
            let input = input_node_ref.cast::<HtmlInputElement>();

            if let Some(input) = input {
                info!("{}", input.value());
            }
        })
    };

    html! {
        <div>
            <label for="target_id">{ "Target ID:" }</label>
            <input ref={input_node_ref} type="text" name="target_id"/>

            <button {onclick}>{ "Create subscription" }</button>
        </div>
    }
}

fn main() {
    wasm_logger::init(wasm_logger::Config::default());
    yew::Renderer::<App>::new().render();
}
