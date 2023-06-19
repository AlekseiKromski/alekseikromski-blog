import BreadCrumbs from "../../../../components/bread-crumbs/bread-crumbs";
import {useSelector} from "react-redux";
import {useEffect, useState} from "react";
import styles from "./genericCU.module.css"
import {useParams} from "react-router-dom";
import {useTranslation} from "react-i18next";

function GenericCU({settings}) {
    let {t} = useTranslation()
    const params = useParams()
    const application = useSelector((state) => state.application);
    const shared = useSelector((state) => state.shared);

    let [data, setData] = useState("")

    useEffect(() => {
        setData(settings.preFunc(params.id, shared, application))
    }, [shared.categories])

    return (
        <div className={styles.createGenericMain}>
            <BreadCrumbs
                breadcrubms={settings.breadcrumbs}
            />

            <label htmlFor="">{t("genericCU.input_name")}</label>
            <input value={data} onChange={(e) => setData(e.target.value)} type="text"/>
            <button onClick={() => {
                if (params.id) {
                    settings.func(params.id, data, application)
                    return
                }
                settings.func(data, application)

            }}>{settings.buttonName}</button>
        </div>
    )
}

export default GenericCU