import React from 'react';
import './Monthly.css';

function Monthly() {
    return <div class="Monthly">
        <h2 class="Monthly-title">2019年12月</h2>
        <table class="Monthly-settlements">
            <tr class="Monthly-settlement">
                <td class="Monthly-settlement-title">商品0</td>
                <td class="Monthly-settlement-amount">1,000</td>
                <td class="Monthly-settlement-amount"></td>
            </tr>
            <tr class="Monthly-settlement">
                <td class="Monthly-settlement-title">商品1</td>
                <td class="Monthly-settlement-amount"></td>
                <td class="Monthly-settlement-amount">2,000</td>
            </tr>
            <tr class="Monthly-settlement Monthly-adjustment">
                <td class="Monthly-settlement-title">月末調整</td>
                <td class="Monthly-settlement-amount">1,000</td>
                <td class="Monthly-settlement-amount">-1,000</td>
            </tr>
            <tr class="Monthly-total">
                <td class="Monthly-settlement-title">合計</td>
                <td class="Monthly-settlement-amount">2,000</td>
                <td class="Monthly-settlement-amount">1,000</td>
            </tr>
        </table>
    </div>;
}

export default Monthly;
