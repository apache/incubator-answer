/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import Progress from './progress';

export default class Confetti {
  private parent: HTMLElement;

  private canvas: HTMLCanvasElement;

  private ctx;

  private width: number;

  private height: number;

  private length: number;

  private yRange: number;

  private progress;

  private rotationRange;

  private speedRange;

  private sprites;

  constructor(param) {
    this.parent = param.elm || document.body;
    this.canvas = document.createElement('canvas');
    this.ctx = this.canvas.getContext('2d');
    this.width = param.width || this.parent.offsetWidth;
    this.height = param.height || this.parent.offsetHeight;
    this.length = param.length || Confetti.CONST.PAPER_LENGTH;
    this.yRange = param.yRange || this.height * 2;
    this.progress = new Progress({
      duration: param.duration,
      isLoop: true,
      timestamp: null,
      delta: 0,
    });
    this.rotationRange =
      typeof param.rotationLength === 'number' ? param.rotationRange : 10;
    this.speedRange =
      typeof param.speedRange === 'number' ? param.speedRange : 10;
    this.sprites = [];

    this.canvas.style.cssText = [
      'display: block',
      'position: absolute',
      'top: 0',
      'left: 0',
      'pointer-events: none',
    ].join(';');

    this.render = this.render.bind(this);

    this.build();

    this.parent.appendChild(this.canvas);
    this.progress.start(performance.now());

    requestAnimationFrame(this.render);
  }

  static get CONST() {
    return {
      SPRITE_WIDTH: 9,
      SPRITE_HEIGHT: 16,
      PAPER_LENGTH: 100,
      DURATION: 8000,
      ROTATION_RATE: 50,
      COLORS: [
        '#EF5350',
        '#EC407A',
        '#AB47BC',
        '#7E57C2',
        '#5C6BC0',
        '#42A5F5',
        '#29B6F6',
        '#26C6DA',
        '#26A69A',
        '#66BB6A',
        '#9CCC65',
        '#D4E157',
        '#FFEE58',
        '#FFCA28',
        '#FFA726',
        '#FF7043',
        '#8D6E63',
        '#BDBDBD',
        '#78909C',
      ],
    };
  }

  build() {
    for (let i = 0; i < this.length; i += 1) {
      const canvas = document.createElement('canvas') as HTMLCanvasElement & {
        position: { initX: number; initY: number };
        rotation;
        speed;
      };
      const ctx = canvas.getContext('2d');

      canvas.width = Confetti.CONST.SPRITE_WIDTH;
      canvas.height = Confetti.CONST.SPRITE_HEIGHT;

      canvas.position = {
        initX: Math.random() * this.width,
        initY: -canvas.height - Math.random() * this.yRange,
      };

      canvas.rotation =
        this.rotationRange / 2 - Math.random() * this.rotationRange;
      canvas.speed =
        this.speedRange / 2 + Math.random() * (this.speedRange / 2);

      if (ctx) {
        ctx.save();
        ctx.fillStyle =
          Confetti.CONST.COLORS[
            Math.floor(Math.random() * Confetti.CONST.COLORS.length)
          ];
        ctx.fillRect(0, 0, canvas.width, canvas.height);
        ctx.restore();
      }

      this.sprites.push(canvas);
    }
  }

  render(now) {
    const progress = this.progress.tick(now);

    this.canvas.width = this.width;
    this.canvas.height = this.height;

    for (let i = 0; i < this.length; i += 1) {
      this.ctx.save();
      this.ctx.translate(
        this.sprites[i].position.initX +
          this.sprites[i].rotation * Confetti.CONST.ROTATION_RATE * progress,
        this.sprites[i].position.initY + progress * (this.height + this.yRange),
      );
      this.ctx.rotate(this.sprites[i].rotation);
      this.ctx.drawImage(
        this.sprites[i],
        (-Confetti.CONST.SPRITE_WIDTH *
          Math.abs(Math.sin(progress * Math.PI * 2 * this.sprites[i].speed))) /
          2,
        -Confetti.CONST.SPRITE_HEIGHT / 2,
        Confetti.CONST.SPRITE_WIDTH *
          Math.abs(Math.sin(progress * Math.PI * 2 * this.sprites[i].speed)),
        Confetti.CONST.SPRITE_HEIGHT,
      );
      this.ctx.restore();
    }

    requestAnimationFrame(this.render);
  }

  destroy() {
    if (this.parent.contains(this.canvas)) {
      this.parent.removeChild(this.canvas);
    }
  }
}
