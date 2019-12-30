import React, { useState } from 'react';
import styles from './NewForm.module.css';

function NewForm() {
        let [enabled, setEnabled] = useState();

        function handleSubmit(event) {
                console.log('submit');
        }

        function toggleEnabled() {
                setEnabled(!enabled);
        }

        return (
                <div className={styles.newForm}>
                        <button className={styles.toggleButton} onClick={toggleEnabled}>New</button>
                        {enabled ? (
                                <form className={styles.form} onSubmit={handleSubmit}>
                                        <dl className={styles.formBody}>
                                                <dt className={styles.inputLabel}>Title</dt>
                                                <dd className={styles.inputField}><input name="inputTitle" className={styles.inputInput}></input></dd>
                                                <dt className={styles.inputLabel}>Amount</dt>
                                                <dd className={styles.inputField}><input name="inputAmount" className={styles.inputInput}></input></dd>
                                        </dl>
                                        <button className={styles.submitButton} type="submit">Submit</button>
                                </form>
                        ) : []}
                </div>
        );
}

export default NewForm;
