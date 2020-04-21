<?php
// automatically generated by the FlatBuffers compiler, do not modify

namespace items;

use \Google\FlatBuffers\Struct;
use \Google\FlatBuffers\Table;
use \Google\FlatBuffers\ByteBuffer;
use \Google\FlatBuffers\FlatBufferBuilder;

class Item extends Table
{
    /**
     * @param ByteBuffer $bb
     * @return Item
     */
    public static function getRootAsItem(ByteBuffer $bb)
    {
        $obj = new Item();
        return ($obj->init($bb->getInt($bb->getPosition()) + $bb->getPosition(), $bb));
    }

    /**
     * @param int $_i offset
     * @param ByteBuffer $_bb
     * @return Item
     **/
    public function init($_i, ByteBuffer $_bb)
    {
        $this->bb_pos = $_i;
        $this->bb = $_bb;
        return $this;
    }

    public function getKey()
    {
        $o = $this->__offset(4);
        return $o != 0 ? $this->__string($o + $this->bb_pos) : null;
    }

    public function getValue()
    {
        $o = $this->__offset(6);
        return $o != 0 ? $this->__string($o + $this->bb_pos) : null;
    }

    public function getTTL()
    {
        $o = $this->__offset(8);
        return $o != 0 ? $this->__string($o + $this->bb_pos) : null;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return void
     */
    public static function startItem(FlatBufferBuilder $builder)
    {
        $builder->StartObject(3);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return Item
     */
    public static function createItem(FlatBufferBuilder $builder, $Key, $Value, $TTL)
    {
        $builder->startObject(3);
        self::addKey($builder, $Key);
        self::addValue($builder, $Value);
        self::addTTL($builder, $TTL);
        $o = $builder->endObject();
        return $o;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param StringOffset
     * @return void
     */
    public static function addKey(FlatBufferBuilder $builder, $Key)
    {
        $builder->addOffsetX(0, $Key, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param StringOffset
     * @return void
     */
    public static function addValue(FlatBufferBuilder $builder, $Value)
    {
        $builder->addOffsetX(1, $Value, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param StringOffset
     * @return void
     */
    public static function addTTL(FlatBufferBuilder $builder, $TTL)
    {
        $builder->addOffsetX(2, $TTL, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return int table offset
     */
    public static function endItem(FlatBufferBuilder $builder)
    {
        $o = $builder->endObject();
        return $o;
    }

    public static function finishItemBuffer(FlatBufferBuilder $builder, $offset)
    {
        $builder->finish($offset);
    }
}
