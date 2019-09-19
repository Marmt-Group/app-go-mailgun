

class MailgunForm extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {

        return (
            <div className="container form-group">
                <form action="/upload" method="post" enctype="multipart/form-data">
                    <div className="field is-grouped">
                        <p className="control">
                            <input className="input" type="text" name="mailgunKey" placeholder="Insert your Mailgun key" />
                        </p>

                        <p className="control">
                            <div className="file has-name">
                                <label className="file-label">
                                    <input className="file-input" type="file" name="file" required />
                                        <span className="file-cta">
                                            <span className="file-icon">
                                                <i className="fas fa-upload"></i>
                                            </span>
                                            <span className="file-label">
                                                Choose a CSV file
                                            </span>
                                        </span>
                                        <span className="file-name">
                                            Screen Shot 2017-07-29 at 15.54.25.png
                                        </span>
                                </label>
                            </div>
                        </p>

                        <p className="control">
                            <input className="button is-primary" type="Submit" value="Submit" />
                        </p>
                    </div>
                </form>
            </div>
        );
    }
}

ReactDOM.render(<MailgunForm />, document.getElementById("mailgun-form"));