import { useEditor, EditorContent } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'

const TiptapView = (props) => {
    const editor = useEditor({
        extensions: [
            StarterKit,
        ],
        content: props.content,
        editable: false
    })

    return (
        <EditorContent editor={editor} />
    )
}

export default TiptapView