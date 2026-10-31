import { Suspense, type FC } from "react";
import { Canvas } from '@react-three/fiber'
import { OrbitControls, useGLTF } from "@react-three/drei";

interface BrainSceneProps{}

const BrainScene: FC<BrainSceneProps> = () => {
    const brainModel = useGLTF('../../public/brain.glb')
    console.log(brainModel)
    return (
        <div className="h-screen w-full bg-[#000000]">
            <Canvas camera={{ position: [0, 0, 5], fov: 50 }}>
                {/* Lights help you see the model if it uses standard materials */}
                <ambientLight intensity={0.5} />
                <directionalLight position={[10, 10, 5]} intensity={1} />
                
                {/* Suspense is required because useGLTF suspends while loading */}
                <Suspense fallback={null}>
                    <primitive object={brainModel.scene} />
                </Suspense>
                
                <OrbitControls />
            </Canvas>
        </div>
    )
}

export default BrainScene