document.addEventListener('DOMContentLoaded', function(){
	const editorContainers = document.querySelectorAll('.editor-container');

	editorContainers.forEach(container => {
		const editorId = container.dataset.editorId;
		const editor = new Quill(container, {
			theme: 'snow'
		});

		editor.on('text-change', function(){
			console.log("text changed")
		});
	});
});

