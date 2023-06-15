import {Link, useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import styles from "./generic.module.css"
import BreadCrumbs from "../../../components/bread-crumbs/bread-crumbs";

function Generic({settings}) {
    const navigate = useNavigate()
    const application = useSelector((state) => state.application);
    let [data, setData] = useState([])

    useEffect(() => {
        //call handler method for fetching data from external API and set it
        settings.data.handler(setData, application)
    }, [])

    function editAction(id) {
        navigate(`${settings.data.urls.edit}/${id}`)
    }

    function deleteAction(id) {
        application.axios.get(`${settings.data.urls.delete}/${id}`)
            .then(() => {
                setData(data.filter(d => {
                    if (d.id !== id) {
                        return d
                    }
                }))
            })
    }

    console.log(data)

    return (
        <div className={styles.posts}>
            <BreadCrumbs
                breadcrubms={settings.breadcrumbs}
            />

            <table>
                <thead>
                    <tr>
                        {
                            settings.data.ids.map(id => (
                                <th>{id}</th>
                            ))
                        }
                    </tr>
                </thead>
                <tbody>
                    {data.length !== 0 &&
                        data.map(d => {
                            return settings.data.renderTable(d, styles, [
                                {
                                    title: "edit",
                                    func: () => {
                                        editAction(d.id)
                                    }
                                },
                                {
                                    title: "delete",
                                    func: () => {
                                        deleteAction(d.id)
                                    }
                                }
                            ])
                        })
                    }
                </tbody>
            </table>
        </div>
    )
}

export default Generic