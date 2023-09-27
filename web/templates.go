package web

import "html/template"

const threadsListHTML string = `
<style>
	body {
		background-color: black;
		color: white;
	}
</style>
<h1 style="text-align:center;">Threads</h1>
<dl>
{{range .Threads}}
	<dt>
		<strong>{{.Title}}<strong>
	</dt>
	<dd>{{.Description}}<dd>
	<dd>
		<form action="/threads/{{.ID}}/delete" method="POST">
			<button type="submit">DELETE</button>
		</form>
	</dd>
{{end}}
</dl>
<a href="/threads/new">Create thread</a>
`

func ThreadsListTemplate() (*template.Template, error) {
	return template.New("List Threads").Parse(threadsListHTML)
}

const threadCreateHTML string = `
<style>
	h2 {
		text-align: center;
	}

	.description, tr {
		display: flex;
		flex-flow: column nowrap;
	}

	.container {
		width: 66vw;
		margin: 1rem auto;
	}

	textarea {
		resize: none
	}
	
</style>

<div class="container">
	<h2>Create New Thread</h2>
	<form action="/threads" method="POST">
		<table>
			<tr>
				<td>Title</td>
				<td><input type="text" name="title" /></td>
			</tr>
		</table>
		<div class="description">
			<label>Description</label>
			<textarea name="description"></textarea>
		</div>
		<button type="submit">Create!</button>
	</form>
</div>
`

func ThreadCreateTemplate() (*template.Template, error) {
	return template.New("Create Thread").Parse(threadCreateHTML)
}
