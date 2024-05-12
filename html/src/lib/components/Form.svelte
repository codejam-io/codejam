<script lang="ts">

    import {setContext} from "svelte";
    import FormField from "./FormField.svelte";
    import {toast} from "svelte-sonner";

    class FormFieldRef {
        component: FormField;
        name: string;
        errorFunction: (msg: string) => {};

        constructor(component: FormField, name: string, errorFunction: (msg: string) => {}) {
            this.component = component;
            this.name = name;
            this.errorFunction = errorFunction;
        }
    }

    let fields : Array<FormFieldRef> = [];

    // Make a registration function available to child FormField components.  FormField components will call
    // this function to register themselves, so this Form can apply error messages to the form based on a
    // FormResponse.
    setContext('register', (fieldComponent, fieldName, errorFunction) => {
            fields.push(new FormFieldRef(fieldComponent, fieldName, errorFunction));
    });

    function SetFieldError(fieldName: string, error: string) {
        fields.forEach((item) => {
            if (item.name === fieldName) {
                item.errorFunction(error);
            }
        })
    }

    export function clearErrors() {
        fields.forEach((item) => {
            item.errorFunction('');
        })
    }

    export function parseResponse(response: Response) {
        if (response.status === 200) {
            toast.success("Event Saved Successfully!");
        } else {
            toast.error("Error Saving.")
            response.json()
                .then((data) => {
                    data.Errors.forEach((error) => {
                        SetFieldError(error.Field, error.Error);
                    })
                })
                .catch((err) => {
                    console.error(err);
                })
        }
    }

</script>

<slot/>
