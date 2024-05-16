<script lang="ts">
    import Page from "../components/Page.svelte";
    import FormField from "../components/FormField.svelte";
    import Form from "../components/Form.svelte";
    import CodejamTeam from "../models/event";
    import { 
        Input, 
        Label, 
        Helper,
        Radio, 
        Textarea,
        MultiSelect } from 'flowbite-svelte';
    
        interface Selections {
            value: string,
            name: string;
        }
    
        let selectedLangs: Selections[] = [];
        let languages: Selections[] = [
            { value: 'py', name: 'Python'},
            { value: 'ph', name: 'PHP'},
            { value: 'js', name: 'JavaScript'},
            { value: 'ts', name: 'TypeScript'}, 
            { value: 'go', name: 'Go'}
        ];
        let languageValues: string[] = languages.map(language => language.value)
    
        
        let selectedTZ: Selections[] = [];
        let listTZ: Selections[] = [
            { value: 'PST', name: 'PST -0700'},
            { value: 'MST', name: 'MST -0600'},
            { value: 'EST', name: 'EST -0400'},
        ]
        let tzValues: string[] = listTZ.map(tz => tz.value)
        let textareaprops = {
            id: 'message',
            name: 'message',
            label: 'What do you want out of this team:',
            rows: 4,
            placeholder: "I want to use threejs so I can learn it!"
        };

let formData : CodeJamTeam | null = null;
let form : Form;
let isSaving : boolean = false;
    
</script>

<Card size="xl" class="w-full">
    <h2>Create your team!</h2>
    {#if formData !== null}
        <div class="flex flex-col gap-8 my-8">
            <Form bind:clearErrors={clearErrors} bind:parseResponse={parseResponse}>
                <FormField name="Status">
                    <Select id="status" items={statusOptions} bind:value={formData.StatusId}></Select>
                </FormField>

                <FormField name="Title">
                    <Input bind:value={formData.Title}></Input>
                </FormField>

                <FormField name="Description">
                    <Textarea rows=10 bind:value={formData.Description}></Textarea>
                </FormField>

                <FormField name="Rules">
                    <Textarea rows=10 bind:value={formData.Rules}></Textarea>
                </FormField>
            </Form>
        </div>

        <Button on:click={saveForm} disabled={isSaving}>
            {#if isSaving}
                <Spinner/>
            {:else}
                Save
            {/if}
        </Button>
    {:else}
        <Spinner/>
    {/if}
</Card>

    <div class="flex flex-col gap-5">
        <Label for="team-name">Team Name</Label>
        <Input id="team-name" placeholder="Team Name:" />

        <br>
        <Radio name="team-type">Public Team</Radio>
        <Helper class="ml-6">(If you want your team to be searchable)</Helper>
        <Radio name="team-type">Private Team</Radio>
        <Helper class="ml-6">(Your team will be invite only)</Helper>

        <MultiSelect id="multi-close" items={languages} bind:value={languageValues} />
        
        <!-- user timezone info will just be displayed for fun -->
        <MultiSelect items={listTZ} bind:value={tzValues} />
        <Label for="tzInput">General Availability:</Label>
        <Input id="tzInput"/>

        <Label for="aboutTextArea">What do you want out of this team?</Label>
        <Textarea {...textareaprops} />

        
    </div>
