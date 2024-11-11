class CubeStatePlayer {
    constructor() {
        this.initializeElements();
        this.initializeState();
        this.addEventListeners();
    }

    initializeElements() {
        // file
        this.fileInput = document.getElementById('fileInput');
        this.fileReset = document.getElementById('fileReset');
        this.fileSelect = document.getElementById('fileSelect');

        // layer
        this.layer1 = document.getElementById('layer1');
        this.layer2 = document.getElementById('layer2');
        this.layer3 = document.getElementById('layer3');
        this.layer4 = document.getElementById('layer4');
        this.layer5 = document.getElementById('layer5')

        // player
        this.startBtn = document.getElementById('startBtn');
        this.prevBtn = document.getElementById('prevBtn');
        this.playPauseBtn = document.getElementById('playPauseBtn');
        this.nextBtn = document.getElementById('nextBtn');
        this.endBtn = document.getElementById('endBtn');

        // speed 
        this.speedSlider = document.getElementById('speedSlider');
        this.progressSlider = document.getElementById('progressSlider');
        this.speedValue = document.getElementById('speedValue');
        this.progressLabel = document.getElementById('progressLabel');
    }

    initializeState() {
        this.states = [];
        this.currentStateIndex = 0;
        this.isPlaying = false;
        this.playbackSpeed = 1.0;
        this.playInterval = null;
    }

    addEventListeners() {
        // file
        this.fileInput.addEventListener('change', () => this.handleFileSelect());
        this.fileReset.addEventListener('click', () => this.handleReset());
        this.fileSelect.addEventListener('change', () => this.loadSelectedFile());


        // player
        this.startBtn.addEventListener('click', () => this.goToStart());
        this.prevBtn.addEventListener('click', () => this.previousState());
        this.playPauseBtn.addEventListener('click', () => this.playPause());
        this.nextBtn.addEventListener('click', () => this.nextState());
        this.endBtn.addEventListener('click', () => this.goToEnd());

        // speed
        this.speedSlider.addEventListener('input', () => this.updateSpeed());
        this.progressSlider.addEventListener('input', () => this.goToPosition());
    }

    // add files to list
    handleFileSelect() {
        const files = Array.from(this.fileInput.files);
        this.updateFileList(files);
    }

    updateFileList(files) {
        files.forEach((file, index) => {
            const option = document.createElement('option');
            option.value = index;
            option.textContent = file.name;
            this.fileSelect.appendChild(option);
        });
    }

    // reset all
    handleReset() {
        this.fileInput.value = '';
        this.fileSelect.innerHTML = '<option value="">Select a file</option>';
        this.states = [];
        this.updateDisplay();
        this.updateControls();
        this.updateSpeed();
    }

    updateDisplay() {
        if (!this.states.length) {
            this.layer1.innerHTML = '';
            this.layer2.innerHTML = '';
            this.layer3.innerHTML = '';
            this.layer4.innerHTML = '';
            this.layer5.innerHTML = '';
            return;
        }

        // read state
        const currentState = this.parseState(this.states[this.currentStateIndex]);
        const previousState = this.currentStateIndex > 0 ? this.parseState(this.states[this.currentStateIndex - 1]) : null;

        this.renderLayer(this.layer1, currentState.layer1, previousState ? previousState.layer1 : []);
        this.renderLayer(this.layer2, currentState.layer2, previousState ? previousState.layer2 : []);
        this.renderLayer(this.layer3, currentState.layer3, previousState ? previousState.layer3 : []);
        this.renderLayer(this.layer4, currentState.layer4, previousState ? previousState.layer4 : []);
        this.renderLayer(this.layer5, currentState.layer5, previousState ? previousState.layer5 : []);

        this.updateProgressLabel();
    }

    parseState(stateStr) {
        // assign state position
        const numbers = stateStr.trim().split(' ').map(Number);
        return {
            layer1: [
                [numbers[0], numbers[1], numbers[2], numbers[3], numbers[4]],
                [numbers[25], numbers[26], numbers[27], numbers[28], numbers[29]],
                [numbers[50], numbers[51], numbers[52], numbers[53], numbers[54]],
                [numbers[75], numbers[76], numbers[77], numbers[78], numbers[79]],
                [numbers[100], numbers[101], numbers[102], numbers[103], numbers[104]],
            ],
            layer2: [
                [numbers[5], numbers[6], numbers[7], numbers[8], numbers[9]],
                [numbers[30], numbers[31], numbers[32], numbers[33], numbers[34]],
                [numbers[55], numbers[56], numbers[57], numbers[58], numbers[59]],
                [numbers[80], numbers[81], numbers[82], numbers[83], numbers[84]],
                [numbers[105], numbers[106], numbers[107], numbers[108], numbers[109]],
            ],
            layer3: [
                [numbers[10], numbers[11], numbers[12], numbers[13], numbers[14]],
                [numbers[35], numbers[36], numbers[37], numbers[38], numbers[39]],
                [numbers[60], numbers[61], numbers[62], numbers[63], numbers[64]],
                [numbers[85], numbers[86], numbers[87], numbers[88], numbers[89]],
                [numbers[110], numbers[111], numbers[112], numbers[113], numbers[114]],
            ],
            layer4: [
                [numbers[15], numbers[16], numbers[17], numbers[18], numbers[19]],
                [numbers[40], numbers[41], numbers[42], numbers[43], numbers[44]],
                [numbers[65], numbers[66], numbers[67], numbers[68], numbers[69]],
                [numbers[90], numbers[91], numbers[92], numbers[93], numbers[94]],
                [numbers[115], numbers[116], numbers[117], numbers[118], numbers[119]],
            ],
            layer5: [
                [numbers[20], numbers[21], numbers[22], numbers[23], numbers[24]],
                [numbers[45], numbers[46], numbers[47], numbers[48], numbers[49]],
                [numbers[70], numbers[71], numbers[72], numbers[73], numbers[74]],
                [numbers[95], numbers[96], numbers[97], numbers[98], numbers[99]],
                [numbers[120], numbers[121], numbers[122], numbers[123], numbers[124]],
            ],

        };
    }

    renderLayer(container, layerData, previousLayerData = []) {
        // update state
        container.innerHTML = '';
        layerData.forEach((row, rowIndex) => {
            row.forEach((value, colIndex) => {
                const cell = document.createElement('div');
                cell.className = 'cube-cell';
                cell.textContent = value;

                if (previousLayerData[rowIndex] && previousLayerData[rowIndex][colIndex] !== value) {
                    cell.classList.add('changed-cell');
                }
                container.appendChild(cell);
            });
        });
    }

    updateProgressLabel() {
        this.progressLabel.textContent =
            `${this.currentStateIndex + 1} / ${this.states.length}`;
    }

    // load selected file
    async loadSelectedFile() {
        const fileIndex = this.fileSelect.value;
        if (fileIndex === '') return;

        const file = this.fileInput.files[fileIndex];
        const content = await file.text();
        this.states = content.trim().split('\n');
        this.currentStateIndex = 0;
        this.updateDisplay();
        this.updateControls();
    }

    updateControls() {
        const hasStates = this.states.length > 0;
        const isFirst = this.currentStateIndex === 0;
        const isLast = this.currentStateIndex === this.states.length - 1;

        this.startBtn.disabled = !hasStates || isFirst;
        this.prevBtn.disabled = !hasStates || isFirst;
        this.playPauseBtn.disabled = !hasStates || isLast;
        this.nextBtn.disabled = !hasStates || isLast;
        this.endBtn.disabled = !hasStates || isLast;

        this.progressSlider.max = Math.max(0, this.states.length - 1);
        this.progressSlider.value = this.currentStateIndex;
    }

    // load initial state
    goToStart() {
        this.currentStateIndex = 0;
        this.updateDisplay();
        this.updateControls();
    }

    // load previous state
    previousState() {
        if (this.currentStateIndex > 0) {
            this.currentStateIndex--;
            this.updateDisplay();
            this.updateControls();
        }
    }

    // trigger play and pause
    playPause() {
        this.isPlaying = !this.isPlaying;
        this.playPauseBtn.innerHTML = this.isPlaying ? '<i class="fa-solid fa-pause"></i>' : '<i class="fa-solid fa-play"></i>';

        if (this.isPlaying) {
            this.startPlayback();
        } else {
            clearInterval(this.playInterval);
        }
    }

    startPlayback() {
        clearInterval(this.playInterval);
        this.playInterval = setInterval(() => {
            if (this.currentStateIndex < this.states.length - 1) {
                this.nextState();
            } else {
                this.isPlaying = false;
                this.playPauseBtn.innerHTML = '<i class="fa-solid fa-play"></i>';
                clearInterval(this.playInterval);
            }
        }, 1000 / this.playbackSpeed);
    }

    // load next state
    nextState() {
        if (this.currentStateIndex < this.states.length - 1) {
            this.currentStateIndex++;
            this.updateDisplay();
            this.updateControls();
        }
    }

    // load final state
    goToEnd() {
        this.currentStateIndex = this.states.length - 1;
        this.updateDisplay();
        this.updateControls();
    }

    // adjust speed
    updateSpeed() {
        this.playbackSpeed = parseFloat(this.speedSlider.value);
        this.speedValue.textContent = this.playbackSpeed.toFixed(1);
        if (this.isPlaying) {
            this.startPlayback();
        }
    }

    // load specific position
    goToPosition() {
        this.currentStateIndex = parseInt(this.progressSlider.value);
        this.updateDisplay();
        this.updateControls();
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const cubeStatePlayer = new CubeStatePlayer();
    cubeStatePlayer.handleReset();
});