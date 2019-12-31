import React, { useState, useEffect } from 'react';
import styles from './Monthly.module.css';
import { useRouteMatch } from 'react-router-dom';
import NewForm from './NewForm';

function Monthly() {
    let { params: { year, month } } = useRouteMatch();
    useEffect(() => {
        fetchRecords();
    }, []);

    let [settlements, setSettlements] = useState([]);
    let [adjustment, setAdjustment] = useState({ amounts: [0, 0] });
    let [total, setTotal] = useState({ amounts: [0, 0] });

    function fetchRecords() {
        setSettlements([
            { title: '商品0', amounts: [1000, 0] },
            { title: '商品1', amounts: [0, 2000] },
        ]);
        setAdjustment({
            amounts: [1000, -1000],
        });
        setTotal({
            amounts: [1000, -1000],
        });
    }

    function number(n) {
        if (n === 0) {
            return "";
        } else {
            return n.toString();
        }
    }

    return (
        <div className="Monthly">
            <header className={styles.header}>
                <h2 className={styles.title}>{year}年{month}月</h2>
                <NewForm />
            </header>
            <table className={styles.settlements}>
                <tbody>
                    {settlements.map((s, i) =>
                        <tr key={i} className={styles.settlement}>
                            <td className={styles.settlementTitle}>{s.title}</td>
                            {s.amounts.map((a, i) =>
                                <td key={i} className={styles.settlementAmount}>{number(a)}</td>
                            )}
                            <td className={styles.deleteColumn}>
                                <button className={styles.deleteButton}>Delete</button>
                            </td>
                        </tr>
                    )}
                    <tr className={[styles.settlement, styles.adjustment].join(' ')}>
                        <td className={styles.settlementTitle}>月末調整</td>
                        {adjustment.amounts.map((a, i) =>
                            <td key={i} className={styles.settlementAmount}>{a}</td>
                        )}
                        <td></td>
                    </tr>
                    <tr className={styles.total}>
                        <td className={styles.settlementTitle}>合計</td>
                        <td className={styles.settlementAmount}>{total.amounts[0]}</td>
                        <td className={styles.settlementAmount}>{total.amounts[1]}</td>
                        <td></td>
                    </tr>
                </tbody>
            </table>
        </div>
    );
}

export default Monthly;
