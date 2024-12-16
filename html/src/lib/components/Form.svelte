<script lang="ts">

    import {setContext} from "svelte";
    import FormField from "./FormField.svelte";
    import {toast} from "svelte-sonner";


    class FormFieldRef {
        name: string;
        errorFunction: (msg: string) => {};

        constructor(name: string, errorFunction: (msg: string) => {}) {
            this.name = name;
            this.errorFunction = errorFunction;
        }
    }

    let fields : Array<FormField> = [];

    // Make a registration function available to child FormField components.  FormField components will call
    // this function to register themselves, so this Form can apply error messages to the form based on a
    // FormResponse.
    setContext('formFunctions', {
        register: (name: string, errorFunction: (msg: string) => {}) => {
            fields.push(new FormFieldRef(name, errorFunction));
        }
    });

    function SetFieldError(fieldName: string, error: string) {
        fields.forEach((item) => {
            if (item.name === fieldName) {
                item.errorFunction(error);
            }
        })
    }

    export const clearErrors = () => {
        fields.forEach((item) => {
            item.errorFunction('');
        })
    }

    export const parseResponse = (response: Response) => {
        if (response.status === 200) {
            toast.success("Save Successful!");
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
