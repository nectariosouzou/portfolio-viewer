import React from 'react';

export class Spinner {
    public elem: HTMLElement | undefined

    start(target: HTMLElement): this {
        this.elem = document.createElement('div')
        target.style.borderRadius = '50%';
        target.style.width = '120px';
        target.style.height = '120px';
        target.style.animation = 'spin 2s linear infinite';
        target.style.border = '16px solid #f3f3f3'; // Light grey
        target.style.borderTop = '16px solid #3498db'; // Blue
        target.appendChild(this.elem)
        return this
    }

    stop() {
        this.elem?.parentNode?.removeChild(this.elem)
        this.elem = undefined
    }
}

