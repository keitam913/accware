import React from 'react';
import './Monthly.css';
import { withRouter } from 'react-router-dom';

class Monthly extends React.Component {
    constructor() {
        super();
        this.state = {
            settlements: [],
            total: { amounts: [0, 0] },
        }
    }

    componentDidMount() {
        let { match } = this.props;
        this.year = match.params.year;
        this.month = match.params.month;
        this.setState({
            settlements: [
                { title: '商品0', amounts: [1000, 0] },
                { title: '商品1', amounts: [0, 2000] },
                { title: '月末調整', amounts: [1000, -1000] },
            ],
        })
        this.setState({
            total: { amounts: [2000, 1000] },
        });
    }

    number(n) {
        if (n === 0) {
            return "";
        } else {
            return n.toString();
        }
    }

    render() {
        return <div className="Monthly">
            <h2 className="Monthly-title">{this.year}年{this.month}月</h2>
            <table className="Monthly-settlements">
                {this.state.settlements.map(s =>
                    <tr className="Monthly-settlement">
                        <td className="Monthly-settlement-title">{s.title}</td>
                        <td className="Monthly-settlement-amount">{this.number(s.amounts[0])}</td>
                        <td className="Monthly-settlement-amount">{this.number(s.amounts[1])}</td>
                    </tr>
                )}
                <tr className="Monthly-total">
                    <td className="Monthly-settlement-title">合計</td>
                    <td className="Monthly-settlement-amount">{this.state.total.amounts[0]}</td>
                    <td className="Monthly-settlement-amount">{this.state.total.amounts[1]}</td>
                </tr>
            </table>
        </div>;
    }
}

export default withRouter(Monthly);
