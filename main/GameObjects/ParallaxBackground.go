components {
  id: "ParallaxBackground"
  component: "/src/sep_scripts/ParallaxBackground.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "bg_03"
  type: "sprite"
  data: "tile_set: \"/Assets/background.atlas\"\n"
  "default_animation: \"BG_03\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.13
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "bg_02"
  type: "sprite"
  data: "tile_set: \"/Assets/background.atlas\"\n"
  "default_animation: \"BG_02\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.14
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "bg_01"
  type: "sprite"
  data: "tile_set: \"/Assets/background.atlas\"\n"
  "default_animation: \"BG_01\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.15
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
