/**
 * @module skottie-slot-manager-sk
 * @description <h2><code>skottie-slot-manager-sk</code></h2>
 *
 * <p>
 *   A component meant to interface with a ManagedAnimation's SlotManager for
 *   property value swapping.
 * </p>
 *
 */

import { ElementSk } from '../../../infra-sk/modules/ElementSk';
import { html } from 'lit-html';
import { SkottieColorEventDetail } from '../skottie-color-input-sk/skottie-color-input-sk';
import { colorToHex, hexToColor } from '../helpers/color';
import { ColorSlot, ScalarSlot, Vec2Slot } from './slot-info';
import { SkottiePlayerSk } from '../skottie-player-sk/skottie-player-sk';
import { define } from '../../../elements-sk/modules/define';
import { SkottieVec2EventDetail } from './skottie-vec2-input-sk';
// without this import, the vec2 input div doesn't populate and vec 2 slots don't render
import './skottie-vec2-input-sk';

export class SkottieSlotManagerSk extends ElementSk {
  private _player: SkottiePlayerSk | null = null;
  private colorSlots: ColorSlot[] = [];
  private scalarSlots: ScalarSlot[] = [];
  private vec2Slots: Vec2Slot[] = [];

  private static template = (ele: SkottieSlotManagerSk) => html`
    <div class="wrapper">
      <ul class="slots-container">
        ${ele.colorSlots.map((item: ColorSlot) => ele.colorSlot(item))}
        ${ele.scalarSlots.map((item: ScalarSlot) => ele.scalarSlot(item))}
        ${ele.vec2Slots.map((item: Vec2Slot) => ele.vec2Slot(item))}
      </ul>
    </div>
  `;

  private colorSlot = (cs: ColorSlot) => html`
    <div class="slot--color">
      <span class="slotID">${cs.id}</span>
      <skottie-color-input-sk
        .color=${cs.colorHex}
        @color-change=${(e: CustomEvent<SkottieColorEventDetail>) =>
          this.onColorSlotChange(e, cs.id)}>
      </skottie-color-input-sk>
    </div>
  `;

  private scalarSlot = (ss: ScalarSlot) => html`
    <div class="slot--scalar">
      <span class="slotID">${ss.id}</span>
      <div class="text-box">
        <input
          type="number"
          class="text-box--input"
          @change=${(e: Event) => this.onScalarSlotChange(e, ss.id)}
          value=${ss.scalar}
          required />
      </div>
    </div>
  `;

  private vec2Slot = (vs: Vec2Slot) => html`
    <div class="slot--vec2">
      <skottie-vec2-input-sk
        .label=${vs.id}
        .x=${vs.x}
        .y=${vs.y}
        @vec2-change=${(e: CustomEvent<SkottieVec2EventDetail>) =>
          this.onVec2SlotChange(e)}>
      </skottie-vec2-input-sk>
    </div>
  `;

  private onColorSlotChange(
    e: CustomEvent<SkottieColorEventDetail>,
    sid: string
  ): void {
    if (!this._player) {
      return;
    }
    const color = hexToColor(e.detail.color);
    this._player.managedAnimation()?.setColorSlot(sid, color);
    this._render();
  }

  private onScalarSlotChange(e: Event, sid: string): void {
    if (!this._player) {
      return;
    }
    const target = e.target as HTMLInputElement;
    this._player.managedAnimation()?.setScalarSlot(sid, Number(target.value));
    this._render();
  }

  private onVec2SlotChange(e: CustomEvent<SkottieVec2EventDetail>): void {
    if (!this._player) {
      return;
    }
    this._player
      .managedAnimation()
      ?.setVec2Slot(e.detail.label, [e.detail.x, e.detail.y]);
  }

  set player(value: SkottiePlayerSk) {
    this._player = value;

    this.colorSlots = [];
    this.scalarSlots = [];
    this.vec2Slots = [];
    const managedAnimation = this._player?.managedAnimation();
    if (managedAnimation) {
      const slotInfo = managedAnimation.getSlotInfo();
      for (const sid of slotInfo.colorSlotIDs) {
        const color = managedAnimation.getColorSlot(sid);
        if (color) {
          const colorHex = colorToHex(Array.from(color));
          this.colorSlots.push({ id: sid, colorHex: colorHex });
        }
      }
      for (const sid of slotInfo.scalarSlotIDs) {
        const scalar = managedAnimation.getScalarSlot(sid);
        if (scalar) {
          this.scalarSlots.push({ id: sid, scalar: scalar });
        }
      }
      for (const sid of slotInfo.vec2SlotIDs) {
        const vec2 = managedAnimation.getVec2Slot(sid);
        if (vec2) {
          this.vec2Slots.push({ id: sid, x: vec2[0], y: vec2[1] });
        }
      }
    }

    this._render();
  }

  constructor() {
    super(SkottieSlotManagerSk.template);
  }

  connectedCallback(): void {
    super.connectedCallback();
    this._render();
  }

  disconnectedCallback(): void {
    super.disconnectedCallback();
  }
}

define('skottie-slot-manager-sk', SkottieSlotManagerSk);
