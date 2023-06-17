import styles from "./alert.module.css"
import CloseIcon from '@mui/icons-material/Close';
import {useEffect, useRef, useState} from "react";
function Alert({title, text, type, set}) {
    let block = useRef()
    let [show, setShow] = useState(false)
    let [listener, setListener] = useState(false)

    useEffect(() => {
        if ( text !== undefined && text != null && text.length !== 0 ){
            setShow(true)
            setTimeout(() => {
                setShow(false)
            }, 5000)
        }
        if (!listener) {
            block.current.addEventListener("animationend", (e) => {
                if (e.target.classList.contains(styles.hide)) {
                    set(null)
                }
            })
            setListener(true)
        }
    }, [text])

    return (
        <div ref={block} className={`${styles.alert} ${
            type === "error" ? styles.error : styles.success
        }
        ${show ? styles.show : styles.hide}
        `}>
            <div className={styles.alertHeader}>
                <h1>{title}</h1>
                <CloseIcon
                    style={{
                        cursor: "pointer"
                    }}
                    onClick={
                        () => {
                            setShow(false)
                        }
                    }
                />
            </div>
            <p>{text}</p>
        </div>
    )
}

export default Alert