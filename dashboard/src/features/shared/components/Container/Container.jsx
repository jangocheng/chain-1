import React from 'react'
import styles from './Container.scss'

class Container extends React.Component {
  render() {
    return (
      <div className={styles.main}>
        {this.props.title && <div className={styles.title}>
          <h5>
            {this.props.title}
          </h5>

          {this.props.actions && <div>{this.props.actions}</div>}
        </div>}

        <div className={styles.children}>
          {this.props.children}
        </div>
      </div>
    )
  }
}

export default Container
